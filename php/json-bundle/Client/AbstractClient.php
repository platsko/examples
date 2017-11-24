<?php

namespace PetrLatsko\JsonBundle\Client;

use PetrLatsko\JsonBundle\Client\Exceptions\ClientErrorJsonException;
use PetrLatsko\JsonBundle\Client\Exceptions\ClientParseJsonException;
use PetrLatsko\JsonBundle\Client\Exceptions\TransportRequestException;
use PetrLatsko\JsonBundle\Client\Interfaces\TransportClientInterface;

abstract class AbstractClient implements TransportClientInterface
{

    /**
     * @var
     */
    protected $dataStore;

    /**
     * @var AbstractTransportProxy
     */
    protected $transport;

    public function getTransport()
    {
        return $this->transport;
    }

    public function setTransport(AbstractTransportProxy $transport)
    {
        $this->transport = $transport;
    }

    /**
     * @param $url
     * @return bool
     * @throws TransportRequestException
     */
    protected function execTransportRequest($url)
    {
        $this->transport->reset();
        $this->transport->execRequest($url);
        $this->transport->prepareErrors();

        return !$this->transport->hasErrors();
    }

    protected function parseJson() {}

    /**
     * @param array $jsonData
     * @return array
     * @throws ClientErrorJsonException
     */
    protected function prepareJsonData()
    {
        if (true !== $this->dataStore['success'])
        {
            throw new ClientErrorJsonException($this->dataStore['message'], $this->dataStore['code']);
        }
    }

    /**
     * @throws ClientParseJsonException
     */
    protected function prepareParseJsonErrors()
    {
        $jsonLastError = json_last_error();
        if (JSON_ERROR_NONE !== $jsonLastError)
        {
            $jsonErrorMessage = 'Unknown error';
            switch ($jsonLastError)
            {
                case JSON_ERROR_DEPTH:
                    $jsonErrorMessage = 'Maximum stack depth exceeded';
                    break;

                case JSON_ERROR_STATE_MISMATCH:
                    $jsonErrorMessage = 'Underflow or the modes mismatch';
                    break;

                case JSON_ERROR_CTRL_CHAR:
                    $jsonErrorMessage = 'Unexpected control character found';
                    break;

                case JSON_ERROR_SYNTAX:
                    $jsonErrorMessage = 'Syntax error, malformed JSON';
                    break;

                case JSON_ERROR_UTF8:
                    $jsonErrorMessage = 'Malformed UTF-8 characters, possibly incorrectly encoded';
                    break;

                default:
                    # Unknown error
                    break;
            }

            throw new ClientParseJsonException($jsonErrorMessage, $jsonLastError);
        }
    }

}
